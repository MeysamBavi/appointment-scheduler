import { Component } from "react";
import Layout from "../components/LayOut";
import Debug from "../components/Debug";
import "../styles/BusinessProfilePage.css";
import {
  Box,
  Grid,
  Paper,
  Table,
  TableBody,
  TableContainer,
  TableHead,
  TableRow,
  TextField,
  TableCell,
  Typography,
  Stack,
  Container,
  List,
  ListItem,
  Button,
} from "@mui/material";

const testData = {
  name: 'پیرایش زیبا',
  kind: 'پیرایشی',
  owner: {
    name: 'رضا',
    family: 'خوش دست',
    phoneNumber: '09123456789'
  }
}



export function PairRow({ rowKey, children }) {
  return (
    <Grid id="top-row" container sx={{ mb: 2 }}>
      <Grid item xs={2} sx={{ mt: 1 }}>
        <Typography variant="item">{rowKey}</Typography>
      </Grid>
      <Grid item xs={10}>
        {children}
      </Grid>

    </Grid >
  );
}

function BusinessProfilePage() {
  return (
    <Layout>
      {/* <Debug text="اینجام"></Debug> */}
      <Container className="todortl" sx={{ px: 1, py: 3 }}>
        <Grid container justifyContent={"space-between"}>
          <Grid item>
            <Typography variant="h5" sx={{ mb: 2 }}>پروفایل</Typography>
          </Grid>
          <Grid item>
            <Button variant="outlined">ذخیره تغییرات</Button>
          </Grid>

        </Grid>


        <hr color="blue"></hr>

        <Box height={20}></Box>


        <PairRow rowKey="نام">
          <TextField fullWidth placeholder="نام" defaultValue={testData.name}></TextField>
        </PairRow>

        <PairRow rowKey="نوع">
          <TextField fullWidth placeholder="نوع" defaultValue={testData.kind}></TextField>
        </PairRow>

        <PairRow rowKey="اطلاعات مالک">
          <Paper variant="outlined" sx={{ px: 1, py: 2 }}>
            <Container>
              <TextField placeholder="نام" defaultValue={testData.owner.name} fullWidth sx={{
                my: 1
              }} />
              <TextField placeholder="نام خانوادگی" defaultValue={testData.owner.family} fullWidth sx={{ my: 1 }} />
              <TextField type="number" placeholder="شماره تلفن" defaultValue={testData.owner.phoneNumber} fullWidth sx={{ my: 1 }} />
            </Container>
          </Paper>
        </PairRow>


      </Container>
    </Layout >
  );
}

export default BusinessProfilePage;
