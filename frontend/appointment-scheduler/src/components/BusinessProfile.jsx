import { Component, useEffect, useState } from "react";
import Layout from "./LayOut";
import Debug from "./Debug";
import "../styles/BusinessProfile.css";
import {
  Box,
  Grid,
  Paper,
  TextField,
  Typography,
  Container,
  Divider,
  Button,
  Autocomplete,
  Stack,
  MenuItem,
  Select,
  useMediaQuery,
} from "@mui/material";
import TimePicker from "react-multi-date-picker/plugins/time_picker";
import { Calendar, DateObject } from "react-multi-date-picker";
import persian from "react-date-object/calendars/persian";
import persian_fa from "react-date-object/locales/persian_fa";
import FormInformationProperty from "./FormInformationProperty";
import {
  readBusiness,
  readBusinessTypes,
  updateBusiness,
} from "../services/ApiService";
import { useLocation } from "react-router-dom";

const testData = {
  ID: 1,
  CreatedAt: "2024-05-25T19:40:42.43623Z",
  UpdatedAt: "2024-05-25T19:40:42.43623Z",
  DeletedAt: null,
  Name: "lksdfjsdf",
  Address: "lsdkfjksdf",
  ServiceType: {
    ID: 1,
    CreatedAt: "2024-05-25T19:26:02.413115Z",
    UpdatedAt: "2024-05-25T19:26:02.413115Z",
    DeletedAt: null,
    Name: "پزشکی",
  },
  ServiceTypeID: 1,
  UserID: 0,
};

// Object { ID: 4, CreatedAt: "2024-05-25T19:57:39.66707Z", UpdatedAt: "2024-05-25T19:57:39.66707Z", DeletedAt: null, Name: "آرایشگاه زیبا", Address: "خوابگاه", ServiceType: {…}, ServiceTypeID: 1, UserID: 0 }

function BusinessProfile() {
  const isDesktop = useMediaQuery("(min-width:600px)");
  const [businessInfo, setBusinessInfo] = useState({});
  const [businessTypes, setBusinessTypes] = useState([]);

  const id = localStorage.getItem("businessId");

  const loadData = async () => {
    const data = await readBusiness(id);
    setBusinessInfo(data);
    console.log(businessInfo);
  };

  const handleBusinessInfoChanges = (e) => {
    setBusinessInfo({
      ...businessInfo,
      [e.target.name]: e.target.value,
    });
  };

  const handleBusinessTypes = async () => {
    const btypes = await readBusinessTypes();
    console.log("something", btypes);
    setBusinessTypes(btypes);
  };

  useEffect(() => {
    // TEST
    // setBusinessesList(testData);
    // ENDTEST

    readBusiness(id).then((data) => setBusinessInfo(data));
    console.log("something");
  }, []);

  return (
    <Container className="todortl" sx={{ px: 1, py: 3 }}>
      <Grid container justifyContent={"space-between"}>
        <Grid item>
          <Typography variant="h5" sx={{ mb: 2 }}>
            پروفایل
          </Typography>
        </Grid>
        <Grid item>
          <Button
            variant="outlined"
            onClick={() => updateBusiness(id, businessInfo)}
          >
            ذخیره تغییرات
          </Button>
          <Button variant="outlined" onClick={() => loadData()}>
            مقادیر قبلی
          </Button>
        </Grid>
      </Grid>

      <Divider />

      <Box height={20}></Box>

      <FormInformationProperty propertyName="نام">
        <TextField
          fullWidth
          name="Name"
          onChange={handleBusinessInfoChanges}
          placeholder="نام"
          value={businessInfo.Name || ""}
        ></TextField>
      </FormInformationProperty>

      <FormInformationProperty propertyName="آدرس">
        <TextField
          fullWidth
          name="Address"
          onChange={handleBusinessInfoChanges}
          placeholder="آدرس"
          value={businessInfo.Address || ""}
        ></TextField>
      </FormInformationProperty>

      {/* TODO: create persian components
        persian needs many settings in react :/ */}
      <FormInformationProperty propertyName="نوع">
        <Autocomplete
          options={businessTypes}
          renderInput={(params) => (
            <TextField
              name="ServiceType"
              value={
                businessInfo.ServiceType ? businessInfo.ServiceType.Name : ""
              }
              onChange={handleBusinessInfoChanges}
              onFocus={handleBusinessTypes}
              {...params}
              label="نوع کسب و کار"
              fullWidth
              sx={{
                "& .MuiAutocomplete-inputRoot": {
                  "& .MuiAutocomplete-input": {
                    textAlign: "right",
                    right: "10px",
                    "& input": {
                      paddingRight: "unset",
                    },
                  },
                },
                "& .MuiAutocomplete-clearIndicator": {
                  marginLeft: 0,
                  marginRight: "0px",
                },
                "& .MuiAutocomplete-popupIndicator": {
                  marginRight: 0,
                },
              }}
              InputLabelProps={{
                sx: {
                  transformOrigin: "right",
                  left: "inherit",
                  right: "1.75rem",
                  fontSize: "small",
                  color: "#807D7B",
                  fontWeight: 400,
                  overflow: "unset",
                },
              }}
            />
          )}
          fullWidth
          getOptionLabel={(option) => option.Name}
        />
      </FormInformationProperty>
    </Container>
  );
}

export default BusinessProfile;
