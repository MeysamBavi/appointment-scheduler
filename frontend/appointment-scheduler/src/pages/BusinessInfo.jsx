import React, { useState } from "react";
import { Sidebar, Menu, MenuItem } from "react-pro-sidebar";
import { FaUser, FaList } from "react-icons/fa";
import Layout from "../components/LayOut";
import WorkersTab from "../components/Table";

const Psidebar = () => {
  const [selectedTab, setSelectedTab] = useState("workers");

  const renderForm = () => {
    switch (selectedTab) {
      case "workers":
        return <WorkersTab />;
      case "appointments":
        return <AppointmentsTab />;
      default:
        return null;
    }
  };

  return (
    <Layout>
      <div style={{ display: "flex", flexDirection: "row-reverse" }}>
        <Sidebar
          breakPoint="md"
          rtl={true}
          collapsed={true} 
          collapsedWidth={0} 
          className="sidebar"
          style={{ flex: "0 0 auto", width: "250px" }}
        >
          <Menu iconShape="circle">
            <MenuItem
              icon={<FaUser />}
              onClick={() => setSelectedTab("workers")}
              style={selectedTab === "workers" ? activeStyle : menuItemStyle}
            >
              کارمندان
            </MenuItem>
            <MenuItem
              icon={<FaList />}
              onClick={() => setSelectedTab("appointments")}
              style={selectedTab === "appointments" ? activeStyle : menuItemStyle}
            >
              نوبت ها
            </MenuItem>
          </Menu>
        </Sidebar>
        <div style={{display: "flex" , justifyContent: "center" , alignItems:"center" , width:"70vw" ,direction:"rtl"}}>{renderForm()}</div>
      </div>
    </Layout>
  );
};

const AppointmentsTab = () => <div>appointmentsTab</div>;

const menuItemStyle = {
  margin: "20px",
};

const activeStyle = {
  ...menuItemStyle,
  border: "1px solid #0076cf",
  borderRadius: "20px",
  backgroundColor : "#0076cf",
  color: "white",
  padding: "5px",
};

export default Psidebar;
