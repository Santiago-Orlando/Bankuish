// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyAlbEP52xQLih-8WpdYAHRP73fmo0MzJ50",
  authDomain: "bankuish-a44ad.firebaseapp.com",
  projectId: "bankuish-a44ad",
  storageBucket: "bankuish-a44ad.appspot.com",
  messagingSenderId: "597403244101",
  appId: "1:597403244101:web:0ce9f7eeba615cb52eb52c",
  measurementId: "G-CNSY7WEJFZ"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const auth = getAuth(app)