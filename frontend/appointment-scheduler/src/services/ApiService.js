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

export const createBusiness = async (businessData) => {
  try {
    const reqbod = {
      name: businessData["businessName"],
      address: businessData["businessAddress"],
      service_type: businessData["businessType"]["ID"],
    };
    const response = await axios.post(apiConfig.userUrl(), reqbod);
    console.log("something: ", axios.defaults.headers.common["Authorization"]);
    console.log(response.data);
  } catch (error) {
    throw new Error(`Error in create business: ${error}`);
  }
};

export const readBusinessTypes = async () => {
  try {
    const response = await axios.get(apiConfig.businessTypeUrl());
    return response.data["service_types"];
  } catch (error) {
    throw new Error(`Error in read business types: ${error}`);
  }
};
