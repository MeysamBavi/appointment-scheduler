import React, { useState } from "react";
import {
  TextField,
  Button,
  Container,
  Stack,
  Autocomplete,
  Stepper,
  Step,
  StepLabel,
  Select,
  MenuItem,
  Typography,
} from "@mui/material";
import Box from "@mui/material/Box";
import { useMediaQuery } from "@mui/material";
import persian from "react-date-object/calendars/persian";
import persian_fa from "react-date-object/locales/persian_fa";
import DatePicker, { DateObject } from "react-multi-date-picker";
import DatePanel from "react-multi-date-picker/plugins/date_panel";
import TimePicker from "react-multi-date-picker/plugins/time_picker";
import { Calendar } from "react-multi-date-picker";

import Layout from "../components/LayOut";
import "../styles/OwnerForm.css";
const OwnerForm = () => {
  const [currentStep, setCurrentStep] = useState(0);
  const [selectedUnit, setSelectedUnit] = useState("min");

  const [ownerInfo, setOwnerInfo] = useState({
    firstName: "",
    lastName: "",
    phoneNumber: "",
  });

  const [businessInfo, setBusinessInfo] = useState({
    businessName: "",
    businessType: null,
  });

  const [appointmentsInfo, setAppointmentsInfo] = useState({
    dateSelection: [new DateObject()],
    hoursSelection: [new DateObject(), new DateObject()],
    appointmentsLength: "",
  });
  const handleOwnerInfoChange = (e) => {
    setOwnerInfo({
      ...ownerInfo,
      [e.target.name]: e.target.value,
    });
  };

  const handleBusinessInfoChange = (e) => {
    setBusinessInfo({
      ...businessInfo,
      [e.target.name]: e.target.value,
    });
  };

  const handleAppointmentsInfoChange = (e) => {
    setAppointmentsInfo({
      ...appointmentsInfo,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log("Owner Info:", ownerInfo);
    console.log("Business Info:", businessInfo);
    console.log("Appointments Info:", appointmentsInfo);
  };
  const isDesktop = useMediaQuery("(min-width:600px)");

  const handleNext = () => {
    setCurrentStep(currentStep + 1);
  };

  const handleBack = () => {
    setCurrentStep(currentStep - 1);
  };

  return (
    <Layout>
      <Container>
        <Box
          sx={{
            boxShadow: "1px 3px 6px 1px #0076cf",
            padding: 3,
            borderRadius: 2,
            maxWidth: 580,
            maxHeight: 520,
            height: 520,
            margin: "auto",
            marginTop: 10,
            marginBottom: 5,
            backgroundColor: "background.paper",
            "& fieldset": {
              borderWidth: "2px !important",
              borderRightWidth: "0px !important",
              borderLeftWidth: "0px !important",
            },
            "& legend": {
              textAlign: "right",
            },
          }}
        >
          <form
            onSubmit={handleSubmit}
            style={{ height: "100%", position: "relative" }}
          >
            <Stepper
              activeStep={currentStep}
              alternativeLabel
              style={{ marginBottom: "15px" }}
            >
              <Step>
                <StepLabel>اطلاعات صاحب کسب و کار</StepLabel>
              </Step>
              <Step>
                <StepLabel>اطلاعات کسب و کار</StepLabel>
              </Step>
              <Step>
                <StepLabel>اطلاعات نوبت ها </StepLabel>
              </Step>
            </Stepper>
            <section style={{ overflow: "auto", height: "calc(100% - 100px)" }}>
              {currentStep === 0 && (
                <Stack spacing={2}>
                  <TextField
                    label="نام"
                    name="firstName"
                    value={ownerInfo.firstName}
                    onChange={handleOwnerInfoChange}
                    fullWidth
                    sx={{
                      "& input": {
                        textAlign: "right",
                      },
                      "& .MuiAutocomplete-inputRoot": {
                        "& .MuiAutocomplete-input": {
                          "& input": {
                            paddingRight: "unset",
                          },
                        },
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
                  <TextField
                    label="نام خانوادگی"
                    name="lastName"
                    value={ownerInfo.lastName}
                    onChange={handleOwnerInfoChange}
                    fullWidth
                    sx={{
                      "& input": {
                        textAlign: "right",
                      },
                      "& .MuiAutocomplete-inputRoot": {
                        "& .MuiAutocomplete-input": {
                          "& input": {
                            paddingRight: "unset",
                          },
                        },
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
                  <TextField
                    label="شماره تماس"
                    name="phoneNumber"
                    value={ownerInfo.phoneNumber}
                    onChange={handleOwnerInfoChange}
                    fullWidth
                    sx={{
                      "& input": {
                        textAlign: "right",
                      },
                      "& .MuiAutocomplete-inputRoot": {
                        "& .MuiAutocomplete-input": {
                          "& input": {
                            paddingRight: "unset",
                          },
                        },
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
              )}

              {currentStep === 1 && (
                <Stack spacing={2}>
                  <TextField
                    label="عنوان کسب و کار"
                    name="businessName"
                    value={businessInfo.businessName}
                    onChange={handleBusinessInfoChange}
                    fullWidth
                    sx={{
                      "& input": {
                        textAlign: "right",
                      },
                      "& .MuiAutocomplete-inputRoot": {
                        "& .MuiAutocomplete-input": {
                          "& input": {
                            paddingRight: "unset",
                          },
                        },
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

                  <Autocomplete
                    options={["املاک", "زیبایی", "سلامت", "موارد دیگر"]}
                    renderInput={(params) => (
                      <TextField
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
                    value={businessInfo.businessType}
                    onChange={(_, value) =>
                      setBusinessInfo({ ...businessInfo, businessType: value })
                    }
                    fullWidth
                    getOptionLabel={(option) => option}
                    isOptionEqualToValue={(option, value) => option === value}
                  />
                </Stack>
              )}

              {currentStep === 2 && (
                <Stack spacing={2}>
                  <Stack
                    direction={"row"}
                    spacing={3}
                    style={
                      isDesktop
                        ? {
                            alignItems: "center",
                            justifyContent: "space-around",
                          }
                        : {
                            alignItems: "center",
                            justifyContent: "space-around",
                            display: "flex",
                            flexDirection: "column",
                          }
                    }
                  >
                    <Calendar
                      editable
                      // multiple="true"
                      range="true"
                      showOtherDays="true"
                      // sort
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

                    <Stack
                      spacing={2}
                      style={
                        isDesktop
                          ? {}
                          : {
                              marginLeft: "0",
                              marginTop: "15px",
                            }
                      }
                    >
                      <Stack spacing={2} textAlign={"right"} fontSize={"small"}>
                        <Typography variant="item">: از ساعت</Typography>
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
                      <Stack spacing={2} textAlign={"right"} fontSize={"small"}>
                        <Typography variant="item">: تا ساعت</Typography>
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
                  <Stack direction={"row"} spacing={2}>
                    <Select
                      value={selectedUnit}
                      onChange={(e) => setSelectedUnit(e.target.value)}
                      variant="outlined"
                      sx={{ width: "30%" }}
                    >
                      <MenuItem value="min">دقیقه</MenuItem>
                      <MenuItem value="hour">ساعت</MenuItem>
                      <MenuItem value="day">روز</MenuItem>
                    </Select>
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
                        },
                        "& .MuiAutocomplete-inputRoot": {
                          "& .MuiAutocomplete-input": {
                            "& input": {
                              paddingRight: "unset",
                            },
                          },
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
              )}
              <Stack
                direction="row"
                spacing={2}
                style={
                  isDesktop || currentStep == 0 || currentStep == 1
                    ? { bottom: "0", position: "absolute" }
                    : { position: "relative", bottom: "0", paddingTop: "20px" }
                }
              >
                {currentStep !== 0 && (
                  <Button variant="contained" onClick={handleBack}>
                    قبلی
                  </Button>
                )}
                {currentStep !== 2 && (
                  <Button variant="contained" onClick={handleNext}>
                    بعدی
                  </Button>
                )}
                {currentStep === 2 && (
                  <Button
                    variant="contained"
                    color="primary"
                    type="submit"
                    onSubmit={handleSubmit}
                  >
                    ارسال
                  </Button>
                )}
              </Stack>
            </section>
          </form>
        </Box>
      </Container>
    </Layout>
  );
};

export default OwnerForm;
