import { useNavigate } from "react-router-dom";
import "./Home.css"

export default function Home() {
  const navigate = useNavigate();
  function checkAuth () {
    navigate("/login");
  }

  return (
    <div className="Home">
      <div className="header"></div>
      <div className="hero"></div>
      <button className="enterLousycord" onClick={checkAuth}>
        Begin with Lousycord
      </button>
    </div>
  )
}
