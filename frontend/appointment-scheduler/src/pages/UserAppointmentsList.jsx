import {
  Container,
  Divider,
  Typography,
  Box,
  Button,
  Grid,
  TextField,
  Paper,
  Link,
  IconButton,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import Layout from "../components/LayOut";
import { useEffect, useState } from "react";
import { deleteBusiness, readBusinesses } from "../services/ApiService";
import { Navigate, useNavigate } from "react-router-dom";

const testData = [
  {
    businessName: "خیاطی",
    businessLink: "https://www.google.com/",
    address: "خیابان سمیه",
    date: "1403.05.01",
    time: "12:00",
  },
  {
    businessName: "پیرایش زیبا",
    businessLink: "https://www.google.com/",
    address: "خیابان بهشتی",
    date: "1403.05.03",
    time: "15:00",
  },
];

function UserAppointmentsList() {
  const [searchString, setSearchString] = useState("");
  const [userAppointments, setUserAppointments] = useState([]);

  const NavigateTo = useNavigate();

  const handleDeleteBusiness = (i) => {
    deleteBusiness(i);
    readBusinesses().then((data) => setUserAppointments(data));
  };

  useEffect(() => {
    // TEST
    setUserAppointments(testData);
    // ENDTEST
    // readBusinesses().then((data) => setUserAppointments(data));
    // console.log("something");
  }, []);

  return (
    <Layout>
      <Container className="todortl" sx={{ px: 1, py: 3 }}>
        <Grid container justifyContent={"space-between"}>
          <Grid item>
            <Typography variant="h5" sx={{ mb: 2 }}>
              نوبت های من
            </Typography>
          </Grid>
          <Grid item>
            <TextField
              sx={{ pb: 2 }}
              fullWidth
              placeholder="جست و جو"
              defaultValue={searchString}
            />
          </Grid>
        </Grid>

        <Divider />

        <Box height={20}></Box>

        {userAppointments.map((item, index) => (
          <Paper
            elevation={3}
            sx={{ height: 50, mb: 1, py: 2, px: 2 }}
            key={index}
          >
            <Grid
              container
              justifyContent={"space-between"}
              alignItems={"center"}
            >
              <Grid item>
                <Grid container>
                  <Link underline="hover" href={item["businessLink"]}>
                    <Typography variant="h6">{item["businessName"]}</Typography>
                  </Link>
                </Grid>
                <Typography>{item["address"]}</Typography>
              </Grid>
              <Grid item>
                <Button variant="outlined">
                  {item["time"]} - {item["date"]}
                </Button>
              </Grid>
            </Grid>
          </Paper>
        ))}
      </Container>
    </Layout>
  );
}

export default UserAppointmentsList;
