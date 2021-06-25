Document is a way to organize and store data as a set of field-value pairs
{
    <field>:<value>
}

Collection - organized store of documents in MongoDB, usually with common fields between documents

Clusters - group of servers that stores your data

Replica Set - a few connected machines that store the same data to ensure that if 
something happens to one of the machines the data will remain intact. Comes from the word replicate - to copy something.

// patch is used to send amll amount of data to the server
// in put request we have to give all the field in body to be updated because it replace whole URI 
// while patch updates only set on instruction