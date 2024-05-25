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

// ----------------------------------------------------- business crud
export const createBusiness = async (businessData) => {
  try {
    const reqbod = {
      name: businessData["businessName"],
      address: businessData["businessAddress"],
      service_type: businessData["businessType"]["ID"],
    };
    const response = await axios.post(apiConfig.businessesListUrl(), reqbod);
    console.log("something: ", axios.defaults.headers.common["Authorization"]);
    console.log(response.data);
  } catch (error) {
    throw new Error(`Error in create business: ${error}`);
  }
};

export const readBusiness = async () => {
  try {
    axios.defaults.headers.common["Authorization"] =
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY2NzEzNzUsImp0aSI6ImQxMTg1MzRjLWJhYzItNGEwOC05ZjI3LTkwMGMxYzdkMGE3OCIsInBob25lX251bWJlciI6IjA5OTA0NjE0MTE2IiwidXNlcl9pZCI6MH0.UTRxysPQuKFzxemVwvB9-9yQb75um8epQYuodNfo3ME";
    const response = await axios.get(apiConfig.businessesListUrl());
    return response.data["businesses"];
  } catch (error) {
    throw new Error(`Error in read business: ${error}`);
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
