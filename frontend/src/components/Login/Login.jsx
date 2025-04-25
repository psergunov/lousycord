import { useState } from "react"
import { Link, useNavigate } from "react-router-dom";
import "./Login.css"

export default function Login() {
  const navigate = useNavigate();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  async function checkAuth() {
    console.log(username, password)
    try {
      const res = await fetch("https://localhost:8080/login", {
        method: "POST",
        headers: {
          "Authorization": "Basic " + btoa(`${username}:${password}`),
          "Content-Type": "application/json",
        },
      });

      if (res.ok) {
        navigate("/");
      } else {
        alert("Unauthorized or invalid credentials.");
      }
    } catch (err) {
      console.error("Auth check failed:", err);
      alert("Failed to check auth.");
    }
  }

  return (
    <div className="Login">
      <div className="loginForm">
        <div className="loginBackground"></div>
        <input type="text" placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)}></input>
        <input type="text" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)}></input>
        <button onClick={checkAuth}>Sign-in</button>
        <Link className="signupLink" to="/signup">
            Or signup instead
        </Link>
      </div>
    </div>
  )
}

