import axios from "axios";
import apiConfig from "./apiConfig";

export const sendOTP = async (phoneNumber) => {
  try {
    const response = await axios.post(apiConfig.otpSendUrl(), {
      phone_number: phoneNumber,
    });
    return response.data;
  } catch (error) {
    throw new Error(`Error sending OTP: ${error}`);
  }
};

export const validateOTP = async (phoneNumber, otp) => {
  try {
    const response = await axios.post(apiConfig.otpValidateUrl(), {
      phone_number: phoneNumber,
      code: otp,
    });
    const jwtToken = response.data.token;
    axios.defaults.headers.common["Authorization"] = `Bearer ${jwtToken}`;
    return jwtToken;
  } catch (error) {
    throw new Error(`Error validating OTP: ${error}`);
  }
};
