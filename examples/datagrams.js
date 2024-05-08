import wt from "k6/x/webtransport";

export default function () {
  const url = "https://localhost:443/webtransport";
  let data = [0, 1, 2, 3, 4];

  wt.connect(url);

  wt.sendDatagram(data);

  const response = wt.receiveDatagram();
  // handle response
}
