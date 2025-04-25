import { useState } from "react"
import { Link, useNavigate } from "react-router-dom";
import "./Signup.css"

export default function Signup() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  return (
    <div className="Signup">
      <div className="signupForm">
        <div className="signupBackground"></div>
        <input type="text" placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)}></input>
        <input type="text" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)}></input>
        <button>Submit</button>
        <Link className="homeLink" to="/">
            Go home
        </Link>
      </div>
    </div>
  )
}