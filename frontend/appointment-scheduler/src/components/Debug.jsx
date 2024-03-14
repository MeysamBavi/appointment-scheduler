import { Box, Container } from "@mui/material";

function Debug(props) {
	return (
		<Box
			height={"100%"}
			width={"100%"}
			display="flex"
			alignItems="center"
			justifyContent={"center"}
			backgroundColor={props.color || "lightblue"}
		>
			{props.text}
		</Box>
	)
}

export default Debug;