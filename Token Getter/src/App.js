import { signInWithEmailAndPassword } from "firebase/auth";
import "./App.css";
import { auth } from "./firebase-config";
import { useState, useEffect } from "react";

function App() {
  const [token, setToken] = useState("");

  const copy = () => {
    navigator.clipboard.writeText(token);
  }

  useEffect( () => {
    (async() => {
      try {
        const user = await signInWithEmailAndPassword(
          auth,
          "seorlando33@gmail.com",
          "12345678"
        );
        const tok = await user.user.getIdToken();
        setToken(tok);
      } catch (error) {
        console.log(error);
      }
    })()
  }, []);

  return (
    <>
    <div className="App">
      <div>
        <h2 className="title">Valid Token:</h2>
        <p>{token}</p>
      </div>
    </div>
    <button onClick={copy}>Copy</button>
    </>
  );
}

export default App;
