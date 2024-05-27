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
    localStorage.setItem("token", jwtToken);
    axios.defaults.headers.common["Authorization"] = `Bearer ${jwtToken}`;
    return jwtToken;
  } catch (error) {
    throw new Error(`Error validating OTP: ${error}`);
  }
};
//GET /businesses/<id:int>/employees
export const getEmployees = async (businessesId) => {
  try {
    const response = await axios.get(apiConfig.getEmployees(businessesId), {
      business_id: businessesId,
    });
    return response;
  } catch (error) {
    throw new Error(`Error getting employees: ${error}`);
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
    const response = await axios.post(apiConfig.businessesListUrl(), reqbod, getHeaders());
    console.log("something: ", jwtToken);
    console.log(response.data);
  } catch (error) {
    throw new Error(`Error in create business: ${error}`);
  }
};

export const readBusinesses = async () => {
  try {
    
    const response = await axios.get(apiConfig.businessesListUrl(), getHeaders());
    return response.data["businesses"];
  } catch (error) {
    throw new Error(`Error in read businesses: ${error}`);
  }
};

export const readBusiness = async (i) => {
  try {
    const response = await axios.get(apiConfig.businessesListUrl() + "/" + i, getHeaders());
    return response.data["business"];
  } catch (error) {
    throw new Error(`Error in read business: ${error}`);
  }
};

export const updateBusiness = async (i, businessData) => {
  try {
    const response = await axios.put(
      apiConfig.businessesListUrl() + "/" + i,
      businessData,
      getHeaders(),
    );
  } catch (error) {
    throw new Error(`Error in update business: ${error}`);
  }
};

export const deleteBusiness = async (i) => {
  try {
    const response = await axios.delete(
      apiConfig.businessesListUrl() + "/" + i, getHeaders(),
    );
  } catch (error) {
    throw new Error(`Error in delete business: ${error}`);
  }
};

// ----------------------------------------------------- business type crud
export const readBusinessTypes = async () => {
  try {
    const response = await axios.get(apiConfig.businessTypeUrl(), getHeaders());
    return response.data["service_types"];
  } catch (error) {
    throw new Error(`Error in read business types: ${error}`);
  }
};

function getHeaders() {
  const jwtToken = localStorage.getItem('token');
  return {
    headers: {
      Authorization: jwtToken,
    }
  }
}
