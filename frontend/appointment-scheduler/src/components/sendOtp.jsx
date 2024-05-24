import React, { useState } from "react";
import { Grid, Button, Typography } from "@mui/material";
import Box from "@mui/material/Box";

import { sendOTP, validateOTP } from "../services/ApiService";
import { useNavigate } from "react-router-dom";
import { MuiOtpInput } from "mui-one-time-password-input";

function SendOTP({ phone , goToStepZero}) {
  const [otp, setOtp] = useState();
  const [otpError, setOtpError] = useState("");
  const [sendingOTP, setSendingOTP] = useState(false);
  const NavigateTo = useNavigate();
  const handleSendOTP = async () => {
    if (!sendingOTP) {
      setSendingOTP(true);
      try {
        await sendOTP(phone);
        alert("OTP sent successfully!");
      } catch (error) {
        alert(`Error sending OTP: ${error.message}`);
      } finally {
        setSendingOTP(false);
      }
    }
  };

  const handleVerifyOTP = async () => {
    if (otp.trim() !== "") {
      try {
        const jwtToken = await validateOTP(phone, otp);
        alert("OTP validated successfully!");
        NavigateTo("/");
      } catch (error) {
        alert(`Error validating OTP: ${error.message}`);
      }
    } else {
      setOtpError("لطفا کد تأیید را وارد کنید");
    }
  };
  const handleChange = (newValue) => {
    setOtp(newValue);
  };
  const handleGoToStepZero = () => {
    goToStepZero(); 
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
        <Typography variant="h6" align="center" gutterBottom>
          کد تأیید را وارد کنید
        </Typography>

        <MuiOtpInput style={{ marginTop: "20px" }} value={otp} onChange={handleChange} />

        <Button
          variant="contained"
          color="primary"
          fullWidth
          onClick={handleVerifyOTP}
          disabled={sendingOTP}
          style={{ marginTop: "20px" }}
        >
          تأیید
        </Button>
        <Button
          variant="contained"
          fullWidth
          onClick={handleSendOTP}
          disabled={sendingOTP}
          style={{ marginTop: "20px" }}
        >
          {sendingOTP ? "در حال ارسال..." : "ارسال مجدد کد"}
        </Button>
        <Typography
          variant="body2"
          color="primary"
          align="right"
          style={{ cursor: "pointer", marginTop: "10px", color: "red" , fontSize: "10px" , textAlign: "right"}}
          onClick={handleGoToStepZero} 
        >
        {`این شماره ${phone} متعلق به شما نیست؟`}
        </Typography>
      </Grid>
    </Box>
  );
}

export default SendOTP;
