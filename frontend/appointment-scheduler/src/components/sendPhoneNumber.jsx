import React, { useState } from "react";
import { Grid, Button, Box, Typography } from "@mui/material";
import TextField from "@mui/material/TextField";
import Sms from "@mui/icons-material/Sms";
import { sendOTP } from "../services/ApiService";
import "../styles/Login.css";
function SendPhoneNumber({ onSend }) {
  const [phone, setPhone] = useState("");
  const [phoneNumberError, setPhoneNumberError] = useState("");
  const [sendingOTP, setSendingOTP] = useState(false);

  const phonePattern = new RegExp(/^(\+98|0)?9\d{9}$/);

  const validatePhone = (e) => {
    setPhone(e.target.value);
    if (phonePattern.test(e.target.value)) {
      setPhoneNumberError("");
    } else {
      setPhoneNumberError("شماره وارد شده اشتباه است!");
    }
  };

  const handleSendOTP = async () => {
    if (phonePattern.test(phone)) {
      
        setSendingOTP(true);
        try {
          await sendOTP(phone);
          onSend(phone);
        } catch (error) {
          alert(error.message);
        } finally {
          setSendingOTP(false);
        }
    } else {
      setPhoneNumberError("!شماره وارد شده اشتباه است");
    }
  };

  return (
    <Box
      display="flex"
      justifyContent="center"
      alignItems="center"
      height="100vh"
    >
      <Grid
        container
        direction="column"
        alignItems="center"
        item
        xs={12}
        md={6}
        sx={{ padding: "20px", fontFamily: "IRANSans !important" }}
      >
        <Grid item>
          <Sms className="otpIcon" />
        </Grid>
        <Typography variant="body1" align="center" gutterBottom>
          شماره موبایل خود را وارد کنید
        </Typography>
        <Typography variant="body2" align="center" gutterBottom>
          یک کد تأیید به شماره تلفنی که شما ارائه می دهید ارسال می شود
        </Typography>
        <div dir="rtl" style={{ width: "100%" }}>
          <TextField
            placeholder="شماره تلفن"
            variant="outlined"
            fullWidth
            value={phone}
            onChange={validatePhone}
            error={phoneNumberError !== ""}
            helperText={phoneNumberError}
            style={{ marginTop: "10px" }}
            type="number"
  inputProps={{ inputMode: "numeric", pattern: "[0-9]*" }}
  onKeyPress={(e) => {
    if (e.key === 'e') {
      e.preventDefault();
    }
  }}
          />
        </div>
        <Button
          variant="contained"
          color="primary"
          fullWidth
          onClick={handleSendOTP}
          disabled={sendingOTP}
          style={{ marginTop: "20px" }}
        >
          {sendingOTP ? "در حال ارسال..." : "ارسال کد"}
        </Button>
      </Grid>
    </Box>
  );
}

export default SendPhoneNumber;
