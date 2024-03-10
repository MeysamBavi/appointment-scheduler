import './App.css'
import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import NotFound from './pages/NotFound';
import Login from './pages/Login';

function App() {
  return (
    <Router>
      <Routes>
        
        {/* route examples: */}
        {/* 
        <Route exact path="/" element={<Bars/>}>
          <Route path="/something" element={<Something><Something/>}/>
          <Route path="/thereis" element={<ThereIs><ThereIs/>}/>
          <Route path="/what" element={<What><What/>}/>
        </Route> */}
        <Route exact path="/login" element={<Login></Login>} />
        <Route path="*" element={<NotFound></NotFound>} />
      </Routes>
    </Router>
  )
}

export default App
