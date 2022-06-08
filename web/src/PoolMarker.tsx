import { ComponentProps } from "react";
import { renderToStaticMarkup } from "react-dom/server";

import { divIcon } from "leaflet";
import { Marker } from "react-leaflet";

import swimmerImage from "./assets/swimmer.png";

const PoolIcon = () => (
  <div
    style={{
      backgroundColor: "black",
      width: "40px",
      height: "40px",
      borderRadius: "25px",
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
    }}
  >
    <img
      style={{ width: "30px", height: "30px", filter: "invert(1)" }}
      src={swimmerImage}
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
