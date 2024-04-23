import { Grid, Button, Hidden, Typography } from "@mui/material";
import landing_desktop from "../assets/landing_desktop.jpg";
import landing_mobile from "../assets/landing_mobile.jpg";
import Layout from "../components/LayOut";
import { Link } from 'react-router-dom';
import "../styles/Landing.css";

function Landing() {
  return (
    <Layout>
      <Grid
        container
        style={{
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
        }}
      >
        <Hidden mdDown>
          <Grid item md={6}>
            <img
              src={landing_mobile}
              alt="landing"
              style={{ width: "100%", height: "auto", objectFit: "cover" }}
            />
          </Grid>
        </Hidden>
        <Grid
          item
          md={6}
          style={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            flexDirection: "column",
          }}
        >
          <Typography
            style={{
              fontFamily: "IRANSans !important",
              fontSize: "1.2em",
              fontWeight: "bold",
              color: "#0076cf",
            }}
            align="center"
            margin={2}
          >
            سرویس نوبت‌دهی آنلاین
          </Typography>
          <Typography
            style={{
              fontFamily: "IRANSans !important",
              fontSize: "0.9em",
              width: "80%",
            }}
            align="center"
            margin={5}
          >
            به سرویس نوبت‌دهی آنلاین خوش آمدید! در اینجا ، افراد و کسب‌وکارها را
            قادر می‌سازیم تا به راحتی خدمات و نوبت‌های خود را مدیریت کنند. اگر
            شما یک ارائه‌دهنده خدمات هستید که می‌خواهید زمان‌های خود را نمایش
            دهید یا یک مشتری که به دنبال راهی ساده برای رزرو وقت است، ما در
            اینجا هستیم تا به شما کمک کنیم
          </Typography>
          <Button
            variant="contained"
            color="primary"
            style={{ marginTop: "20px", height: "max-content" }}
            component={Link}
            to="/owner-appointmets"
          >
            ساخت صفحه شخصی
          </Button>
        </Grid>
      </Grid>
    </Layout>
  );
}

export default Landing;
