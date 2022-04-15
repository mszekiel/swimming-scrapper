import fetch from "node-fetch";

const getData = () => {
  fetch(
    "https://functions.api.ticos-systems.cloud/api/gates/counter?organizationUnitIds=30208",
    {
      headers: {
        "Abp-TenantId": 69,
        "Abp.TenantId": 69,
        Origin: "https://www.swm.de",
        Referer: "https://www.swm.de/",
      },
    }
  ).then((response) =>
    db.addMessage(JSON.stringify({ date: new Date(), ...response.data[0] }))
  );
};
