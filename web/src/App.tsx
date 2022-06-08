import styled from "@emotion/styled";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";

import "leaflet/dist/leaflet.css";
import { useQuery } from "react-query";
import { getPools } from "./api/pools";
import { useEffect } from "react";
import { PoolMarker } from "./PoolMarker";

const Container = styled.div`
  width: 100%;
  height: 100%;
`;

function App() {
  const { data } = useQuery("pools", getPools);

  useEffect(() => {
    console.log(data?.data?.map((i) => [i.location.lat, i.location.lon]));
  }, [data]);

  return (
    <Container>
      <MapContainer
        style={{ width: "100%", height: "100%" }}
        center={[48.137154, 11.576124]}
        zoom={13}
        scrollWheelZoom={true}
      >
        <TileLayer
          attribution='Map tiles by <a href="http://stamen.com">Stamen Design</a>, <a href="http://creativecommons.org/licenses/by/3.0">CC BY 3.0</a> &mdash; Map data &copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
          url="https://stamen-tiles-{s}.a.ssl.fastly.net/toner-lite/{z}/{x}/{y}{r}.png"
        />
        {data?.data?.map((i) => (
          <PoolMarker position={[i.location.lat, i.location.lon]} />
        ))}
      </MapContainer>
    </Container>
  );
}

export default App;
