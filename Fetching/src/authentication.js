'use strict';

const fs = require('fs');
const jwt = require(`jsonwebtoken`);
const filePath = `../data.json`

module.exports = {
    jwtVerify: async (req, res, next) => {
        const file = JSON.parse(fs.readFileSync(filePath,`utf8`))
        const token = req.headers['x-access-token'];
        if (!token) return res.status(401).send({ auth: false, message: 'No token provided.' });
  
        jwt.verify(token, file.secretKey, function(err, decoded) {
          if (err) return res.status(500).send({ auth: false, message: 'Failed to authenticate token.' });
        req.user = decoded;
        return next();
    })
}
}