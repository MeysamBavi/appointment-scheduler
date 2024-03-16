import React, { useState } from "react";
import { Grid,useMediaQuery  } from "@mui/material";
import otp_recieve from "../assets/otp_recieve.jpg";
import SendPhoneNumber from "../components/sendPhoneNumber"
import SendOTP from "../components/sendOtp";

function Login() {
  const [step, setStep] = useState(0);
  const [phone, setPhone] = useState("");
  const isDesktop = useMediaQuery("(min-width:900px)");

  const handleSendPhoneNumber = (phone) => {
    setPhone(phone);
    setStep(1); 
  };
  const handleGoToStepZero = () => {
    setStep(0);
  };
  return (
    <Grid container  overflow="hidden">
      {isDesktop && (
        <Grid item xs={12} md={6} sx={{ height: "100vh" }}>
          <img
            src={otp_recieve}
            alt="otp"
            style={{ width: "100%", height: "100%", objectFit: "cover" }}
          />
        </Grid>
      )}
      <Grid item xs={12} md={6}>
        {step === 0 ? (
          <SendPhoneNumber onSend={handleSendPhoneNumber} />
        ) : (
          <SendOTP phone={phone} goToStepZero={handleGoToStepZero}/>
        )}
      </Grid>
    </Grid>
  );
}

export default Login;
