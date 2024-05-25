import { Component, useState } from "react";
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

function BusinessProfile() {
  const isDesktop = useMediaQuery("(min-width:600px)");
  const [selectedUnit, setSelectedUnit] = useState("min");
  const [businessInfo, setBusinessInfo] = useState({
    name: "پیرایش زیبا",
    businessType: "زیبایی",
    owner: {
      name: "رضا",
      family: "خوش دست",
      phoneNumber: "09123456789",
    },
  });

  const [appointmentsInfo, setAppointmentsInfo] = useState({
    dateSelection: [new DateObject()],
    hoursSelection: [new DateObject(), new DateObject()],
    appointmentsLength: "",
  });

  const handleAppointmentsInfoChange = (e) => {
    setAppointmentsInfo({
      ...appointmentsInfo,
      [e.target.name]: e.target.value,
    });
  };

  return (
    <Layout>
      <Container className="todortl" sx={{ px: 1, py: 3 }}>
        <Grid container justifyContent={"space-between"}>
          <Grid item>
            <Typography variant="h5" sx={{ mb: 2 }}>
              پروفایل
            </Typography>
          </Grid>
          <Grid item>
            <Button variant="outlined">ذخیره تغییرات</Button>
          </Grid>
        </Grid>

        <Divider />

        <Box height={20}></Box>

        <FormInformationProperty propertyName="نام">
          <TextField
            fullWidth
            placeholder="نام"
            defaultValue={businessInfo.name}
          ></TextField>
        </FormInformationProperty>

        {/* TODO: create persian components
        persian needs many settings in react :/ */}
        <FormInformationProperty propertyName="نوع">
          <Autocomplete
            options={["املاک", "زیبایی", "سلامت", "موارد دیگر"]}
            renderInput={(params) => (
              <TextField
                {...params}
                InputProps={{
                  ...params.InputProps,
                  endAdornment: (
                    <div>{params.InputProps.endAdornment.props.children}</div>
                  ),
                  startAdornment: null,
                }}
                placeholder="نوع کسب و کار"
                fullWidth
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
        </FormInformationProperty>

        <FormInformationProperty propertyName="اطلاعات مالک">
          <Paper variant="outlined" sx={{ px: 1, py: 2 }}>
            <Container>
              <TextField
                placeholder="نام"
                defaultValue={businessInfo.owner.name}
                fullWidth
                sx={{
                  my: 1,
                }}
              />
              <TextField
                placeholder="نام خانوادگی"
                defaultValue={businessInfo.owner.family}
                fullWidth
                sx={{ my: 1 }}
              />
              <TextField
                type="number"
                placeholder="شماره تلفن"
                defaultValue={businessInfo.owner.phoneNumber}
                fullWidth
                sx={{ my: 1 }}
              />
            </Container>
          </Paper>
        </FormInformationProperty>

        {/* TODO: make calendar and times as one component
        this is just copy & paste */}
        <FormInformationProperty propertyName="اطلاعات نوبت ها">
          <Paper variant="outlined" sx={{ px: 1, py: 2 }}>
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
                    <label variant="h6">: از ساعت</label>
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
                    <label variant="h6">: تا ساعت</label>
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
              </Stack>
            </Stack>
          </Paper>
        </FormInformationProperty>
      </Container>
    </Layout>
  );
}

export default BusinessProfile;
