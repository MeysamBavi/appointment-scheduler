import {
    Grid,
    Typography,
} from "@mui/material";

export default function FormInformationProperty({ propertyName, children }) {
    return (
        <Grid id="top-row" container sx={{ mb: 2 }}>
            <Grid item xs={2} sx={{ mt: 1 }}>
                <Typography variant="item">{propertyName}</Typography>
            </Grid>
            <Grid item xs={10}>
                {children}
            </Grid>
        </Grid >
    );
}