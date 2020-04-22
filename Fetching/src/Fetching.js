'use strict';

const axios = require(`axios`);
const fs = require('fs');
const conversionRate = {
    "IdrToUsd": 0.000065041767,
    "apiKey": "6b2df0571955f97c3e28",
    "apiUrl": "https://free.currconv.com/api/v7/convert",
    "updatedAt": 1587335161991
    }

async function conversion(file) {
    const timeGap = (new Date().getTime() / 1000) - conversionRate.updatedAt;
    if (timeGap < 3600 * 6) return conversionRate.IdrToUsd;

    const rate = await axios({
        method: `GET`,
        url: `${conversionRate.apiUrl}`,
        params:{
            q: `IDR_USD`,
            compact: `ultra`,
            apiKey: conversionRate.apiKey
        }
    }).catch(function(error) {
        console.log(error);
        throw(error);
    });
    conversionRate.IdrToUsd = rate.data.IDR_USD;
    conversionRate.updatedAt = new Date().getTime() / 1000;

    return rate.data.IDR_USD;
}

function calculate(array) {
	let sum	= 0;
	let median = 0;
	for (let i=0; i < array.length; i++) {
		for (let j=0; j < array.length; j++) {
			if (array[j] > array[j+1]) {
				let temp = array[j]
				array[j] = array[j + 1]
				array[j+1] = temp
			}
		}
	}
	
	for (let i=0; i < array.length; i++) {
		sum += array[i]
	}
	
	if (array.length % 2 === 0) {
		median = (array[array.length / 2 - 1] + array[array.length / 2]) / 2
	}	else {
		median = array[(array.length - 1)/2]
	}
	
	const result = {
		max : array[array.length - 1],
		min: array[0],
		median,
		avg: sum / array.length
	}
	return result
}


function time(a) {
    let now 
    if (a.length < 13) now = new Date(parseInt(a) * 1000)
    else now = new Date(parseInt(a))
	const firstDayOfYear = new Date(now.getFullYear(), 0, 1);
    let pastDaysOfYear = (now - firstDayOfYear) / 86400000;
	const week =  Math.ceil((pastDaysOfYear + firstDayOfYear.getDay() + 1) / 7);
	const year = now.getFullYear();
	const result = `week ${week} of ${year}`
    return result
}


module.exports = {
    index: async(req, res) => {
        const convRate = await conversion(conversionRate);

        const fetchData = await axios({
            method: `GET`,
            url: `https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list`
        }).catch(function(error) {
            console.log(error);
            throw(error);
        })

        let resultData = fetchData.data
        for (let i = 0; i < resultData.length; i++) {
            const priceUsd = parseInt(resultData[i].price) * convRate;
            resultData[i].price_usd = `${priceUsd}`
        }

        return res.status(200).json({result: resultData});
    },

    aggregate: async(req, res) => {
        try {
            if (req.user.role !== `admin`) return res.status(401).json({error: `admin access only`});
            const fetchData = await axios({
                method: `GET`,
                url: `https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list`
            }).catch(function(error) {
                console.log(error);
                throw(error);
            })
            
            let temp =[];
		
            for (let i = 0; i < fetchData.data.length; i++) {
                if (fetchData.data[i].area_provinsi && fetchData.data[i].price) {
                    temp.push({
                        provinsi : fetchData.data[i].area_provinsi,
                        price : parseInt(fetchData.data[i].price),
                        time : time(fetchData.data[i].timestamp)
                    })
                }
            }
            
            let temp1 = [];
            for (let i = 0; i < temp.length; i++) {
                if (temp1.length <= 0) {
                    temp1.push({
                        provinsi: temp[i].provinsi,
                        price:[[temp[i].price]],
                        time: [temp[i].time]
                    })
                }
                    const indexProvinsi = temp1.map(function (e) {
                        return e.provinsi
                    }).indexOf(temp[i].provinsi)
                    if (indexProvinsi < 0) {
                        temp1.push({
                            provinsi: temp[i].provinsi,
                            price:[[temp[i].price]],
                            time: [temp[i].time]
                        })
                    } else {
                        const indexTime = temp1[indexProvinsi].time.indexOf(temp[i].time)
                        if (indexTime < 0) {
                            temp1[indexProvinsi].time.push(temp[i].time)
                            temp1[indexProvinsi].price.push([temp[i].price])
                        } else {
                            temp1[indexProvinsi].price[indexTime].push(temp[i].price)
                        }
                    }
            }
            
            let result = []
            let resultObj ={}
            for (let i=0; i < temp1.length; i++) {
            for (let j = 0; j < temp1[i].time.length; j++) {
                if (result.length <= 0) {
                    resultObj[`province`] = temp1[i].provinsi
                    resultObj[temp1[i].time[j]] = calculate(temp1[i].price[j])
                    result.push(resultObj)
                } else {
                    const index = result.map(function (e) {return e.province}).indexOf(temp1[i].provinsi)
                    if (index < 0) {
                        resultObj[`province`] = temp1[i].provinsi
                        resultObj[temp1[i].time[j]] = calculate(temp1[i].price[j])
                        result.push(resultObj)
                    } else {
                    result[index][[temp1[i].time[j]]] = calculate(temp1[i].price[j])
                    }
                }
            }
            resultObj = {}
            }

            return res.status(200).json({data: result});
        } catch (error) {
            console.log(error);
        }
    },

    retrieve: async(req, res) => {
        try {
            return res.status(200).json({claim: req.user});
        } catch (error) {
            console.log(error);
        }
    }
}