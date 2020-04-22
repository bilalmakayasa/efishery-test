'use strict';

const express = require('express');
const app = express();
const http = require('http');
const Fetching = require(`./src/Fetching`);
const auth = require(`./src/authentication`);


app.set('port', 8082);
app.get('/', function(req, res){
  res.send('Hello World');
});
app.get(`/testing`, auth.jwtVerify, Fetching.index);
app.get(`/aggregate`, auth.jwtVerify, Fetching.aggregate);
app.get(`/retrieve`, auth.jwtVerify, Fetching.retrieve);

http.createServer(app).listen(app.get('port'), function(){
    console.log("Fetching services listening on port " + app.get('port'));
  });