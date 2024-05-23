import { Component } from "react";
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
} from "@mui/material";
import FormInformationProperty from "../components/FormInformationProperty";

const testData = {
  name: "پیرایش زیبا",
  kind: "پیرایشی",
  owner: {
    name: "رضا",
    family: "خوش دست",
    phoneNumber: "09123456789",
  },
};

function BusinessProfile() {
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
            defaultValue={testData.name}
          ></TextField>
        </FormInformationProperty>

        <FormInformationProperty propertyName="نوع">
          <TextField
            fullWidth
            placeholder="نوع"
            defaultValue={testData.kind}
          ></TextField>
        </FormInformationProperty>

        <FormInformationProperty propertyName="اطلاعات مالک">
          <Paper variant="outlined" sx={{ px: 1, py: 2 }}>
            <Container>
              <TextField
                placeholder="نام"
                defaultValue={testData.owner.name}
                fullWidth
                sx={{
                  my: 1,
                }}
              />
              <TextField
                placeholder="نام خانوادگی"
                defaultValue={testData.owner.family}
                fullWidth
                sx={{ my: 1 }}
              />
              <TextField
                type="number"
                placeholder="شماره تلفن"
                defaultValue={testData.owner.phoneNumber}
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
