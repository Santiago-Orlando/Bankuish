import { signInWithEmailAndPassword } from "firebase/auth";
import "./App.css";
import { auth } from "./firebase-config";
import { useState } from "react"

function App() {
  const [token, setToken] = useState("");

  const [emailLogin, setEmailLogin] = useState("");
  const [passwordLogin, setPasswordLogin] = useState("");

  const login = async () => {
    try {
      const user = await signInWithEmailAndPassword(auth, emailLogin, passwordLogin)
      const tok = await user.user.getIdToken()
      console.log(tok);
    } catch (error) {
      console.log(error);
    }
  }

  return (
    <div className="App">
      <input
        type="email"
        onChange={(event) => {
          setEmailLogin(event.target.value);
        }}
      />
      <input
        type="password"
        onChange={(event) => {
          setPasswordLogin(event.target.value);
        }}
      />
      <button onClick={login}>login</button>
    </div>
  );
}

export default App;
