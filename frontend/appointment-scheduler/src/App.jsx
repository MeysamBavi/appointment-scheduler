import "./styles/App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import NotFound from "./pages/NotFound";
import Login from "./pages/Login";
import Landing from "./pages/Landing";
import OwnerForm from "./pages/OwnerForm";
import BusinessProfilePage from "./pages/BusinessProfilePage";

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
        <Route exact path="/" element={<Landing></Landing>} />
        <Route exact path="/login" element={<Login></Login>} />
        <Route
          exact
          path="/ownerCreateForm"
          element={<OwnerForm></OwnerForm>}
        />
        <Route
          exact
          path="/business-profile"
          element={<BusinessProfilePage />}
        />
        <Route path="*" element={<NotFound></NotFound>} />
      </Routes>
    </Router>
  );
}

export default App;
