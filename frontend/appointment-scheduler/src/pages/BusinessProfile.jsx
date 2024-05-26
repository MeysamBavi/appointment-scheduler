import React, { useEffect, useState } from "react";
import Layout from "../components/LayOut";
import Debug from "../components/Debug";
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
import FormInformationProperty from "../components/FormInformationProperty";
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

function BusinessProfile() {
  const isDesktop = useMediaQuery("(min-width:600px)");
  const [businessInfo, setBusinessInfo] = useState({});
  const [businessTypes, setBusinessTypes] = useState([]);
  const { state } = useLocation();
  const id = state["id"];

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

  return (
    <Layout>
      <Container className="todortl" sx={{ px: 1, py: 3 }}>
        <Grid container justifyContent="space-between">
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
                      my: 1,
                    },
                  },
                }}
              />
            )}
            fullWidth
            getOptionLabel={(option) => option.Name}
          />
        </FormInformationProperty>

        <FormInformationProperty propertyName="اطلاعات نوبت ها">
          <Paper variant="outlined" sx={{ px: 1, py: 2 }}>
            <Stack spacing={2}>
              <Stack
                direction={isDesktop ? "row" : "column"}
                spacing={3}
                alignItems="center"
                justifyContent={isDesktop ? "space-around" : "flex-start"}
              >
                <Calendar
                  editable
                  range
                  showOtherDays
                  value={appointmentsInfo.dateSelection}
                  placeholder="تاریخ شروع رویداد"
                  format="YYYY/MM/DD"
                  onChange={(value) =>
                    setAppointmentsInfo({
                      ...appointmentsInfo,
                      dateSelection: value,
                    })
                  }
                  rangeHover
                  calendar={persian}
                  locale={persian_fa}
                />

                <Stack spacing={2} style={{ marginLeft: isDesktop ? 0 : "0px", marginTop: isDesktop ? 0 : "15px" }}>
                  <Stack spacing={2} textAlign="right" fontSize="small">
                    <Typography>: از ساعت</Typography>
                    <Calendar
                      disableDayPicker
                      format="HH:mm A"
                      plugins={[<TimePicker hideSeconds />]}
                      value={appointmentsInfo.hoursSelection[0]}
                      onChange={(_, value) =>
                        setAppointmentsInfo({
                          ...appointmentsInfo,
                          hoursSelection: [
                            value,
                            appointmentsInfo.hoursSelection[1],
                          ],
                        })
                      }
                    />
                  </Stack>
                  <Stack spacing={2} textAlign="right" fontSize="small">
                    <Typography>: تا ساعت</Typography>
                    <Calendar
                      disableDayPicker
                      format="HH:mm A"
                      plugins={[<TimePicker hideSeconds />]}
                      value={appointmentsInfo.hoursSelection[1]}
                      onChange={(_, value) =>
                        setAppointmentsInfo({
                          ...appointmentsInfo,
                          hoursSelection: [
                            appointmentsInfo.hoursSelection[0],
                            value,
                          ],
                        })
                      }
                    />
                  </Stack>
                </Stack>
              </Stack>
              <Stack direction="row" spacing={2}>
                <TextField
                  label={`مدت هر نوبت`}
                  type="number"
                  name="appointmentsLength"
                  value={appointmentsInfo.appointmentsLength}
                  onChange={handleAppointmentsInfoChange}
                  fullWidth
                  sx={{
                    "& input": {
                      textAlign: "right",
                      right: "10px",
                      "& input": {
                        paddingRight: "unset",
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
              </Stack>
            </Stack>
          </Paper>
        </FormInformationProperty>
      </Container>
    </Layout>
  );
}

export default BusinessProfile;
