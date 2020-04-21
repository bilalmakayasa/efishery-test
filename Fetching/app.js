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
app.get(`/aggregate`, Fetching.aggregate);
app.get(`/retrieve`, auth.jwtVerify, Fetching.retrieve);

http.createServer(app).listen(app.get('port'), function(){
    console.log("Express server listening on port " + app.get('port'));
  });