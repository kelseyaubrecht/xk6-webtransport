import wt from "k6/x/webtransport";

export default function () {
  const url = "https://localhost:443/webtransport";
  wt.connect(url);

  let data = [0, 1, 2, 3, 4];
  wt.write(data);

  const response = wt.readAll();
  // handle response

  wt.close();
}
