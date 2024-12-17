import './App.css';
import { Link } from "react-router-dom";
function Header(){
    return(
    <div className="header">
        <h1>Go Study</h1>
        <nav>
        <Link to="/">Home</Link>
        <Link to="/login">Login</Link>
        <Link to="/signup">Signup</Link>
        </nav>
      </div>)
}
export default Header;