import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./components/Home/Home.jsx"
import Login from "./components/Login/Login.jsx"
import Signup from "./components/Signup/Signup.jsx"

function App() {

  return (
    <>
      <Router>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="/" element={<Home />} />
        </Routes>  
      </Router>
    </>
  )
}

export default App
