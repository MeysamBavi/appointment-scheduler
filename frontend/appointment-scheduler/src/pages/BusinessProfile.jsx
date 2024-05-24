import { Component, useState } from "react";
import Layout from "../components/LayOut";
import Debug from "../components/Debug";
import "../styles/BusinessProfilePage.css";
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
} from "@mui/material";
import FormInformationProperty from "../components/FormInformationProperty";

function BusinessProfile() {
  const [businessInfo, setBusinessInfo] = useState({
    name: "پیرایش زیبا",
    businessType: "زیبایی",
    owner: {
      name: "رضا",
      family: "خوش دست",
      phoneNumber: "09123456789",
    },
  });

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
                label="نوع کسب و کار"
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
      </Container>
    </Layout>
  );
}

export default BusinessProfile;
