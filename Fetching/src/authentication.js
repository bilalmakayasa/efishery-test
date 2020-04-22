'use strict';

const fs = require('fs');
const jwt = require(`jsonwebtoken`);
const filePath = `./data.json`

module.exports = {
    jwtVerify: async (req, res, next) => {
        const file = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQClSEK7AtfQnRmDFCFvdbM6D6F+tP8eKOWqnvs0jnInOvrZ4c0Nzu/9ilzEhYjYMh+P79hEy0jw1+09/u6lrYo6/BAjxIpSOZDiOZlFKylUnK60fxc6auJv4NjyOBKoVXKqQUuCju6QjdYRTl+6XrSms1/fpEXFMRJP4IMthjO2DT3d6oNuzQhScAleDl9lsOYvcitn3+uZttHdVXjYsU/luw1/OPjMqf2gh7cR+GX7lX1nCa9RYAYijpYNT6LAA/XQUSYGl9qVfFR8ZaWYyTQgITUKQQHac3GpOhBBFyrkJYdNG+L34072f/Jj3qe7eKDD1SR5szyU+/S3aFd4YB+/Z2YAD+0YLud09NoIlHbisrgoYsPlyz2r+Hri0tW7IyDKhjyymbGVEBC3WKHTCW1LpOR5B/+u9nNm5AVqZ76GMOjbXyWDaKz40FnwToIpDe5oxhq/fMINbXrPViivtjfGT3ifFrzl+9j4U0qSSUAxmUE88yRubzekmEbqX++3KUJtlfclpiPAP4HB2ayVBDSUGMhSrqmmrLxggB8nTnX61jFxkBOCiN+LyOeRbwf+3RkTQ5yhu+/OBP2L6msWqTyhO77FlZGzktHqfexqspbLTjWcmMJFvt6tJgFdk3IIXAnLeCEbuoDcl2C60Y/jrrsDFeCbtvzuA4GSMtNjUs76rQ== bilal.makayasa@gmail.com"

        const token = req.headers['x-access-token'];
        if (!token) return res.status(401).send({ auth: false, message: 'No token provided.' });
  
        jwt.verify(token, file, function(err, decoded) {
          if (err) return res.status(500).send({ auth: false, message: 'Failed to authenticate token.' });
        req.user = decoded;
        return next();
    })
}
}