import express, { Request, Response } from "express";
import cors from "cors";
import dotenv from "dotenv";
import mongoose, { ConnectOptions, MongooseOptions } from "mongoose";
import { ussdRoute } from "./ussd.route";
import { TokenFactoryClient } from "./shared/services/blockchain/blockchain-client/index";
import { BeepTxClient } from "./shared/services/blockchain/blockchain-client/tx";

(async () => {
  // const tokenFactoryClient = new TokenFactoryClient('http://0.0.0.0:26657')
  // const connectany = await tokenFactoryClient.createAccount()
  // // const balance = await tokenFactoryClient.getNativeTokenBal("beep1v928q8czt84flwdzw8nfm2lwyse32n8kguj575")

  // const connectWallet = await tokenFactoryClient.connecWallet("city crawl text view all hybrid fee output crush horror foil receive")

  // const beepTxClient =  new BeepTxClient(connectWallet.client)

  // console.log('account', connectany)
  // // console.log('balance', balance)
})()


dotenv.config();

const app = express();
const PORT = process.env.PORT || 5000;

// Middleware
app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

// Sample Route
app.get("/", (req: Request, res: Response) => {
  res.send("Hello, Express with TypeScript!");
});

const MONGODB_URI = process.env.MONGODB_URI as string;
(async () => {
    try {
      mongoose.connect(MONGODB_URI, {
        useNewUrlParser: true,
        useUnifiedTopology: true,
      } as ConnectOptions);
      console.log("Connected To Database - Initial Connection");
    } catch (err) {
      console.log(
        `Initial Distribution API Database connection error occurred -`,
        err
      );
    }
})();

app.post('/ussd', ussdRoute) 

// app.post('/ussd', (req, res) => {
//   // Read the variables sent via POST from our API
//   const {
//       sessionId,
//       serviceCode,
//       phoneNumber,
//       text,
//   } = req.body;

//   let response = '';

//   if (text == '') {
//       // This is the first request. Note how we start the response with CON
//       console.log('sessionId', sessionId)
//       console.log('serviceCode', serviceCode)
//       console.log('phoneNumber', phoneNumber)
//       console.log('text', text)
//       response = `CON What would you like to check
//       1. My account
//       2. My phone number
//       2. My wife Name`;
//   } else if ( text == '1') {
//       // Business logic for first level response
//       response = `CON Choose account information you want to view
//       1. Account number`;
//   } else if ( text == '2') {
//       // Business logic for first level response
//       // This is a terminal request. Note how we start the response with END
//       response = `END Your phone number is ${phoneNumber}`;
//   } else if ( text == '1*1') {
//       // This is a second level response where the user selected 1 in the first instance
//       const accountNumber = 'ACC100101';
//       // This is a terminal request. Note how we start the response with END
//       response = `END Your account number is ${accountNumber}`;
//   }else if (text == '3'){
//       response = `END Your wife name is Zainab`;
//   }

//   // Send the response back to the API
//   res.set('Content-Type: text/plain');
//   res.send(response);
// });

// Start Server
app.listen(PORT, () => {
  console.log(`🚀 Server is running on http://localhost:${PORT}`);
});