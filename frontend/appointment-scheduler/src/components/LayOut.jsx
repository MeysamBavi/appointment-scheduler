import React, { useState } from "react";
import { AppBar, Toolbar, IconButton, Menu, MenuItem, ClickAwayListener } from "@mui/material";
import MenuIcon from '@mui/icons-material/Menu';
import Button from '@mui/material/Button';
import { Link } from 'react-router-dom';
import '../styles/LayOut.css';

function Layout({ children }) {
  const [anchorEl, setAnchorEl] = useState(null);

  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <div style={{ overflow: "hidden", position: "fixed" }}>
      <AppBar position="static">
        <Toolbar>
          <ClickAwayListener onClickAway={handleClose}>
            <div>
              <IconButton
                size="large"
                edge="start"
                color="inherit"
                aria-label="menu"
                aria-controls="menu"
                aria-haspopup="true"
                onClick={handleClick}
                sx={{ mr: 2 }}
              >
                <MenuIcon />
              </IconButton>
              <Menu
                id="menu"
                anchorEl={anchorEl}
                open={Boolean(anchorEl)}
                onClose={handleClose}
              >
                <MenuItem
                  component={Link}
                  to="/customer-appointments"
                  onClick={handleClose}
                  className="menu-item"
                >
                  نوبت های من
                </MenuItem>
                <MenuItem
                  component={Link}
                  to="/owner-appointmets"
                  onClick={handleClose}
                  className="menu-item"
                >
                  صفحات من
                </MenuItem>
              </Menu>
            </div>
          </ClickAwayListener>
          <Button color="inherit" component={Link} to="/login">ورود</Button>
        </Toolbar>
      </AppBar>
      <div>{children}</div>
    </div>
  );
}

export default Layout;
