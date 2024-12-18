package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"

)
const mongoURI = "mongodb://localhost:27017"
var client *mongo.Client


func home(w http.ResponseWriter,r *http.Request){
fmt.Fprintln(w)
fmt.Fprintln(w,"hello")
}




func main(){
	fmt.Println("server not yet started....")

	http.HandleFunc("/",home)
	connectdb()
	createcollection()
	http.HandleFunc("/getdata",getdata)
	http.HandleFunc("/postdata",postdata)
	http.HandleFunc("/deletedata",deletedata)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		fmt.Println("server failed to start",err)
	}

}
func postdata(w http.ResponseWriter,r *http.Request){
	if r.Method!=http.MethodPost{
		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type","application/json")
	var newData map[string]interface{}
	err:=json.NewDecoder(r.Body).Decode(&newData)
	if err!=nil{
		http.Error(w,"invalid json format",http.StatusBadRequest)
		return
	}
	fmt.Println("new data reciveed:",newData)
	dbname:="tnr"
	colname:="student"
	collection:=client.Database(dbname).Collection(colname)
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	result,err:=collection.InsertOne(ctx,newData)
	if err!=nil{
		http.Error(w,"error ocuured while inserting",http.StatusInternalServerError)
		return
	}
	fmt.Println("data inserted scueesfully",result.InsertedID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":"data inserted successfully",
		"id":result.InsertedID,
	})
}
func getdata(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w, "that Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	dbname:="tnr"
	colname:="student"
	collection:=client.Database(dbname).Collection(colname)
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	cursor,err:=collection.Find(ctx,bson.D{})
	if err!=nil{
		http.Error(w,"error occured",http.StatusInternalServerError)
		return
	}
	var results []bson.M
	err=cursor.All(ctx,&results)
	if err!=nil{
		http.Error(w,"error occured",http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	err=json.NewEncoder(w).Encode(results)
	if err!=nil{
		http.Error(w,"error occured",http.StatusInternalServerError)
		return
	}

}
func deletedata(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	dbname := "tnr"
	colname := "student"
	collection := client.Database(dbname).Collection(colname)

	// Step 1: Parse the JSON body dynamically
	var filter bson.M
	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Step 2: Check if the filter is empty
	if len(filter) == 0 {
		http.Error(w, "No filter provided", http.StatusBadRequest)
		return
	}

	// Step 3: Delete the document based on the filter
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		http.Error(w, "Error while deleting data", http.StatusInternalServerError)
		return
	}

	// Step 4: Check if the document was actually deleted
	if result.DeletedCount == 0 {
		fmt.Fprintf(w, "❌ No records found matching filter: %v", filter)
	} else {
		fmt.Fprintf(w, "✅ Data deleted successfully matching filter: %v", filter)
	}
}

func createcollection(){
	dbname:="tnr"
	colname:="student"
	collection:=client.Database(dbname).Collection(colname)
	/*dummyDocument := bson.D{
		{Key: "name", Value: "TNR"},
		{Key: "age", Value: 22},
		{Key:"course",Value:"CSE"},
		{Key:"Collage",Value:"SRM"},
	}
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	result,err:=collection.InsertOne(ctx,dummyDocument)
	if err!=nil{
		fmt.Printf("failed")
	}else{
		fmt.Printf("successfully done %v",result.InsertedID)
	}*/
	if collection!=nil{
		fmt.Println("successfully connected")
	} else{
		fmt.Println("error in connection")
	}
	
}

func connectdb(){
	// Create a context with a 10-second timeout for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()


	var err error
	clientoptions:= options.Client().ApplyURI(mongoURI)
	client,err=mongo.Connect(ctx,clientoptions)
	if err!=nil{
		fmt.Println("failed to connect mongodb:",err)
		return
	}
	err=client.Ping(ctx,nil)
	if err!=nil{
		fmt.Println("Failed to ping mongodb: ",err)
		return
	}
	fmt.Println("connected to mongodb successfully")
}