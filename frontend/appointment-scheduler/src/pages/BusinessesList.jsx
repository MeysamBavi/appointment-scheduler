import {
  Container,
  Divider,
  Typography,
  Box,
  Button,
  Grid,
  TextField,
  Paper,
  SvgIcon,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import Layout from "../components/LayOut";
import { useEffect, useState } from "react";
import Icon from "react-multi-date-picker/components/icon";
import Edit from "@mui/icons-material/Edit";
import { readBusiness } from "../services/ApiService";

const testData = [
  {
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
  },
  {
    ID: 2,
    CreatedAt: "2024-05-25T19:51:47.989891Z",
    UpdatedAt: "2024-05-25T19:51:47.989891Z",
    DeletedAt: null,
    Name: "arayesh",
    Address: "tehran",
    ServiceType: {
      ID: 2,
      CreatedAt: "2024-05-25T19:26:02.421769Z",
      UpdatedAt: "2024-05-25T19:26:02.421769Z",
      DeletedAt: null,
      Name: "آرایشی",
    },
    ServiceTypeID: 2,
    UserID: 0,
  },
  {
    ID: 3,
    CreatedAt: "2024-05-25T19:55:13.340435Z",
    UpdatedAt: "2024-05-25T19:55:13.340435Z",
    DeletedAt: null,
    Name: "آرایشگاه زیبا",
    Address: "خوابگاه",
    ServiceType: {
      ID: 1,
      CreatedAt: "2024-05-25T19:26:02.413115Z",
      UpdatedAt: "2024-05-25T19:26:02.413115Z",
      DeletedAt: null,
      Name: "پزشکی",
    },
    ServiceTypeID: 1,
    UserID: 0,
  },
  {
    ID: 4,
    CreatedAt: "2024-05-25T19:57:39.66707Z",
    UpdatedAt: "2024-05-25T19:57:39.66707Z",
    DeletedAt: null,
    Name: "آرایشگاه زیبا",
    Address: "خوابگاه",
    ServiceType: {
      ID: 1,
      CreatedAt: "2024-05-25T19:26:02.413115Z",
      UpdatedAt: "2024-05-25T19:26:02.413115Z",
      DeletedAt: null,
      Name: "پزشکی",
    },
    ServiceTypeID: 1,
    UserID: 0,
  },
  {
    ID: 6,
    CreatedAt: "2024-05-25T20:51:10.768144Z",
    UpdatedAt: "2024-05-25T20:51:10.768144Z",
    DeletedAt: null,
    Name: "خیاطی",
    Address: "خیاطخونه",
    ServiceType: {
      ID: 1,
      CreatedAt: "2024-05-25T19:26:02.413115Z",
      UpdatedAt: "2024-05-25T19:26:02.413115Z",
      DeletedAt: null,
      Name: "پزشکی",
    },
    ServiceTypeID: 1,
    UserID: 0,
  },
];

function BusinessesList() {
  const [searchString, setSearchString] = useState("");
  const [businessesList, setBusinessesList] = useState([]);

  //   const handleBusinessesList = () => {
  //     const rbs = readBusiness();
  //     setBusinessesList(rbs);
  //   };

  useEffect(() => {
    readBusiness().then((data) => setBusinessesList(data));
  });

  return (
    <Layout>
      <Container className="todortl" sx={{ px: 1, py: 3 }}>
        <Grid container justifyContent={"space-between"}>
          <Grid item>
            <Typography variant="h5" sx={{ mb: 2 }}>
              کسب و کار های من
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

        {businessesList.map((item, index) => (
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
                  <Typography variant="h6">{item["Name"]}</Typography>
                  <Button disabled>{item["ServiceType"]["Name"]}</Button>
                </Grid>
                <Typography>{item["Address"]}</Typography>
              </Grid>
              <Grid item>
                <EditIcon color="primary" sx={{ ml: 1 }}></EditIcon>
                <DeleteIcon color="error"></DeleteIcon>
              </Grid>
            </Grid>
          </Paper>
        ))}
      </Container>
    </Layout>
  );
}

export default BusinessesList;
