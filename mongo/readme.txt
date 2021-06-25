username: moo1-student
password: m001-mongodb-basics
cluster name: M001
database name: m001-student

// connect to atlas:
mongo "mongodb+srv://<username>:<password>@<cluster>.mongodb.net/admin"
mongo "mongodb+srv://moo1-student:m001-mongodb-basics@M001.mongodb.net/admin"
mongo "mongodb+srv://sandbox.fko8v.mongodb.net/m001-student" --username m001-student

// list database
show dbs

// go into database
use sample_training


// display collection
show collections

// find command
db.zips.find({"state": "NY"})
db.zips.find({"state": "NY"}).count()
db.zips.find({"state": "NY", "city": "ALBANY"})
db.zips.find({"state": "NY", "city": "ALBANY"}).pretty()
db.inspections.findOne() // pick up a random document from the database

// insert command
db.inspections.insert({
      "_id" : ObjectId("56d61033a378eccde8a8354f"),
      "id" : "10021-2015-ENFO",
      "certificate_number" : 9278806,
      "business_name" : "ATLIXCO DELI GROCERY INC.",
      "date" : "Feb 20 2015",
      "result" : "No Violation Issued",
      "sector" : "Cigarette Retail Dealer - 127",
      "address" : {
              "city" : "RIDGEWOOD",
              "zip" : 11385,
              "street" : "MENAHAN ST",
              "number" : 1712
         }
  })

// without object id 
db.inspections.insert({
      "id" : "10021-2015-ENFO",
      "certificate_number" : 9278806,
      "business_name" : "ATLIXCO DELI GROCERY INC.",
      "date" : "Feb 20 2015",
      "result" : "No Violation Issued",
      "sector" : "Cigarette Retail Dealer - 127",
      "address" : {
              "city" : "RIDGEWOOD",
              "zip" : 11385,
              "street" : "MENAHAN ST",
              "number" : 1712
         }
  })

It command iterate through the cursor results


// how does _id assigned to a document??
It is automatically generated as an ObjectId type value.
You can select a non ObjectId type value when inserting a new document, 
as long as that value is unique to this document.


// to avoid duplication
You can place additional rules on which documents can and 
cannot be inserted into a collection using MongoDB's schema validation functionality.

// insert multiple data
db.inspections.insert([ { "test": 1 }, { "test": 2 }, { "test": 3 } ])
db.inspections.insert([{ "_id": 1, "test": 1 },{ "_id": 1, "test": 2 },
                       { "_id": 3, "test": 3 }],{ "ordered": false })


// update operations
db.zips.updateMany({ "city": "HUDSON" }, { "$inc": { "pop": 10 } })
db.zips.updateOne({ "zip": "12534" }, { "$set": { "pop": 17630 } })
db.grades.updateOne({ "student_id": 250, "class_id": 339 },
                    { "$push": { "scores": { "type": "extra credit",
                                             "score": 100 }
                                }
                     })

// delete operations
db.inspections.deleteMany({ "test": 1 })
db.<collection name>.drop()


// comparation operator
$eq, $nq, $lt, $gt, $lte, $gte
db.trips.find({ "tripduration": { "$lte" : 70 },
                "usertype": { "$ne": "Subscriber" } }).pretty()


// logical operators
$and, $or, $nor, $not
{"student_id": {"$gt": 25, "$lte": 26}}
db.routes.find({ "$and": [ { "$or" :[ { "dst_airport": "KZN" },
                                    { "src_airport": "KZN" }
                                  ] },
                          { "$or" :[ { "airplane": "CR2" },
                                     { "airplane": "A81" } ] }
                         ]}).pretty()

db.companies.find(
    {
        $and : [
            { "$or": [
                { "founded_year": 2004 },
                { "founded_month": 10 }
            ]},
            { "$or": [
                { "category_code": "social" },
                { "category_code": "web" }
            ]}
        ]
    }
)
.count()

// expressive operator
$expr
db.trips.find({
    "$expr": {
        "$and": [
            { "$gt": ["$tripduration", 1200] },
            { "$eq": ["$end station id", "$start station id"] }
        ]
    }
}).count()

// Array operator
$push: It allows us an element to an array
$size: Returns a cursor with all documents where the specified array field is the exactly given length.
{ <array field> : { "$size" : <number> }}
{ <array field> : { "$all" : <array> }}
{ <array field> : [element] } // exact match
$all: return a cursor with all documetns in which the specified array field contains all the given 
elements regardless of their order in the array.

{ "accommodates": {"$gt": 6}, "reviews": { "$size": 50 } }
{ "accommodates": {"$gt": 6}, "reviews": { "$size": 50 }, {"name": 1} } // shows only specific field
db.listingsAndReviews.find({ "property_type": "House", "amenities": "Changing table" }).count()

// projection syntax
db.<collection>.find({ <query> }, { <projection> } )
1 - include the field
0 - exclude the field
Use only 1s or only 0s

// sub element of array
db.grades.find({ "class_id":431 }, {"scores":{$elemMatch:{"score":{"$gt":85}}}}).count()

// array of sub document
db.trips.findOne({ "start station location.type": "Point" }) 
{ "address.city": {"$regex": "NEW YORK"} } // regex is to compare exact pattern of string

// Aggregation Framework
$match: will check for matching string
$project: will displays only specific fields
$group: it will count total number of items
db.listingsAndReviews.aggregate([
                    { "$match": { "amenities": "Wifi" } },
                    { "$project": { "price": 1, "address": 1, "_id": 0 }}]).pretty()

db.listingsAndReviews.aggregate([
                { "$project": { "address": 1, "_id": 0 }},
                { "$group":   { "_id": "$address.country", "count": { "$sum": 1 } } }])

// sort and init method
db.zips.find().sort({"pop":-1}).limit(1).pretty()
1 for  ascending
-1 for descending

// index
db.trips.find({ "birth year": 1989 })
db.trips.find({ "start station id": 476 }).sort( { "birth year": 1 } )
db.trips.createIndex({ "birth year": 1 })
db.trips.createIndex({ "start station id": 1, "birth year": 1 })

// data modeling
data is stored in the way that it is used.
data is accessed together stored together

// upsert
db.collection.updateOne({<query>}, {<update>}, {"upsert":true})
update will happen if given criteria is matched
insert will happen if given criteria is not matched
db.iot.updateOne({ "sensor": r.sensor, "date": r.date,
                   "valcount": { "$lt": 48 } },
                    { "$push": { "readings": { "v": r.value, "t": r.time } },
                    "$inc": { "valcount": 1, "total": r.value } },
                 { "upsert": true })

