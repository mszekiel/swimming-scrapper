import { ComponentProps } from "react";
import { Marker } from "react-leaflet";
import { renderToStaticMarkup, renderToString } from "react-dom/server";
import { divIcon } from "leaflet";
import { css } from "@emotion/react";

const PoolIcon = () => (
  <div
    style={{
      backgroundColor: "black",
      width: "50px",
      height: "50px",
      borderRadius: "25px",
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
    }}
  >
    <img
      style={{ width: "35px", height: "35px", filter: "invert(1)" }}
      src="assets/swimmer.png"
    />
  </div>
);

export const PoolMarker = ({ position }: ComponentProps<typeof Marker>) => {
  const poolIcon = divIcon({
    html: renderToStaticMarkup(<PoolIcon />),
    className: "custom-icon",
    iconSize: [50, 50],
    iconAnchor: [25, 47.5],
  });

  return <Marker position={position} icon={poolIcon} />;
};
