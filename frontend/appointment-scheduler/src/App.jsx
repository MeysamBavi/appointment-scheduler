import "./styles/App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import NotFound from "./pages/NotFound";
import Login from "./pages/Login";
import Landing from "./pages/Landing";
import OwnerForm from "./pages/OwnerForm";
import AppSidebar from "./pages/BusinessInfo";
import BusinessesList from "./pages/BusinessesList";
import UserAppointmentsList from "./pages/UserAppointmentsList";

function App() {
  return (
    <Router>
      <Routes>
        <Route exact path="/" element={<Landing></Landing>} />
        <Route exact path="/login" element={<Login></Login>} />
        <Route
          exact
          path="/ownerCreateForm"
          element={<OwnerForm></OwnerForm>}
        />
        <Route exact path="/businesses-list" element={<BusinessesList />} />
        <Route exact path="/business-info" element={<AppSidebar />} />
        <Route
          exact
          path="/user-appointments-list"
          element={<UserAppointmentsList />}
        />

        <Route path="*" element={<NotFound></NotFound>} />
      </Routes>
    </Router>
  );
}

export default App;
