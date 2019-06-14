const request = require('request-promise');

export function test() {
    return 'Caicaicai'
}

export async function createDir() {

    let options;
    let result;
    let directoryID;
  
    /*
      Create a new directory
      POST /directory/new/
      STEP1: apply and obtain a new directory ID
    */
    console.log('Create a new directory');
    try {
      options = {
        url: 'https://dxdl.deepq.com:5000/directory/new/',
        form: {}
      };
      await request.post(options, (error, res, body) => {
        result = JSON.parse(res.body).result;
        directoryID = result.directoryID;
        console.log(result);
      });
    } catch (error) {
      console.log(error.message);
    }
    return directoryID;
  }

  export async function regProvider( directoryID , userType , providerID , providerPassword ) {
    let options;
    let result;
  /*
    Register a user (provider)
    POST /user/register/
    STEP2: apply an user account on the new directory
           enter the directory ID from STEP1, user type, user-specified ID and password
  */
  console.log('Register a provider');
  try {
    options = {
      url: 'https://dxdl.deepq.com:5000/user/register/',
      form: {
        directoryID: directoryID,
        userType: userType,
        userID: providerID,
        password: providerPassword
      }
    };
    await request.post(options, (error, res, body) => {
      result = JSON.parse(res.body).result;
      console.log(result);
    });
  } catch (error) {
    console.log(error.message);
  }
  return result.message;
}

export async function regConsumer( directoryID , consumerID , consumerPassword ) {

    let options;
    let result;
  /*
    Register a user (consumer)
    POST /user/register/
    STEP3: apply an user account on the new directory
           enter the directory ID from STEP1, user type, user-specified ID and password
  */
  console.log('Register a consumer');
  try {
    options = {
      url: 'https://dxdl.deepq.com:5000/user/register/',
      form: {
        directoryID: directoryID,
        userType: 'consumer',
        userID: consumerID ,
        password: consumerPassword
      }
    };
    await request.post(options, (error, res, body) => {
      result = JSON.parse(res.body).result;
      console.log(result);
    });
  } catch (error) {
    console.log(error.message);
  }
}

export async function createDataEntry( directoryID , 
  providerID , 
  providerPassword,
  dataOwner,
  dataDescription,
  dataCertificate,
  dataAccessPath,
  offerPrice,
  dueDate ) {


  let options;
  let result;
/*
  Create a new data entry by a provider
  POST /entry/create/
  STEP4: create a new data entry
         enter the directory ID from STEP1, the provider ID and password from STEP2
         data entry = data owner, data description, data certificate, data access path,
                      offer price and data entry due date
*/
console.log('Create a new data entry by a provider');
try {
  options = {
    url: 'https://dxdl.deepq.com:5000/entry/create/',
    form: {
      directoryID: directoryID,
      userID: providerID,
      password: providerPassword,
      dataOwner: dataOwner,
      dataDescription: dataDescription ,
      dataCertificate: dataCertificate ,
      dataAccessPath: dataAccessPath ,
      offerPrice: offerPrice ,
      dueDate: dueDate 
    }
  };
  await request.post(options, (error, res, body) => {
    result = JSON.parse(res.body).result;
    console.log(result);
  });
} catch (error) {
  console.log(error.message);
}

  return result.message ;
}

export async function retrieveEAS( EASID ) {
  let options;
  let result;
/*
  Retrieve EAS information by an EAS ID
  GET /eas/sid/
  STEP10: retrieve EAS information by an EAS ID
          enter the EAS ID from STEP5
          If the the API /eas/revoke/ is called --> isValid=false
          If the the API /eas/revoke/ is not called --> isValid=true
*/ 
console.log('Retrieve EAS information by an EAS ID');
try {
  options = {
    url: 'https://dxdl.deepq.com:5000/eas/sid/' + '?EASID=' + EASID,
  };
  await request.get(options, (error, res, body) => {
    result = JSON.parse(res.body).result;
    console.log(result);
  });
} catch (error) {
  console.log(error.message);
}
  return result.dataAccessPath;
}


export async function EAS( directoryID , 
  consumerID,
  dataCertificate,
  expirationDate,
  providerAgreement,
  consumerAgreement ) {

  let options;
  let result;
  let EASID;
/*
  Deploy a new EAS for a provider and a consumer by the service
  POST /eas/deploy/
  STEP5: deploy a new EAS and returned an EAS ID
         enter the directory ID from STEP1 and the consumer ID from STEP2
         EAS = data entry fields + 
               data certificate, data expiration date, provider agreement and consumer agreement
*/
console.log('Deploy a new EAS for a provider and a consumer by the service');
try {
  options = {
    url: 'https://dxdl.deepq.com:5000/eas/deploy/',
    form: {
      directoryID: directoryID,
      userID: consumerID,
      dataCertificate: dataCertificate,
      expirationDate: expirationDate,
      providerAgreement: providerAgreement,
      consumerAgreement: consumerAgreement
    }
  };
  await request.post(options, (error, res, body) => {
    result = JSON.parse(res.body).result;
    EASID = result.EASID;
    console.log(EASID);
    console.log(result);
  });
} catch (error) {
  console.log(error.message);
}
  return EASID;
}