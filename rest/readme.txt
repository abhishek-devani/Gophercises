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

CORS - cross origin resourse sharing
// same origin : 2 URLs have same origins if they have identical schemes, host and ports.
// unless the request is from same origin, the request will be rejected by browser.
// same origin means same scheme, host or port
//CORS is a standard to relax the same-origin policy.

// how to enable
// add config.enableCORS in WebAPIConfig
// Add [EnableCors] attribute to the controller class.

X-Requested-With. 
Mainly used to identify Ajax requests 
(most JavaScript frameworks send this field with value of XMLHttpRequest ); 
also identifies Android apps using WebView. 
X-Requested-With: XMLHttpRequest. DNT. 
Requests a web application to disable their tracking of a user.