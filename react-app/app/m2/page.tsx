// App.tsx
"use client";
import React, { useState, ChangeEvent } from "react";
import { useSearchParams } from "next/navigation";
import { Line } from "react-chartjs-2";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";

// Register the required Chart.js components
ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

const App: React.FC = () => {
  let tcpData = require("../[data]/m2/GeeToPi.json");
  const searchParams = useSearchParams();

  let file = searchParams.get("file");
  if (file === "GeeToPi") {
    tcpData = require("../[data]/m2/GeeToPi.json");
  } else if (file == "localToGee") {
    tcpData = require("../[data]/m2/localToGee.json");
  } else if (file == "localTolocal") {
    tcpData = require("../[data]/m2/localTolocal.json");
  } else {
    file = "localToPi";
    tcpData = require("../[data]/m2/localToPi.json");
  }

  // Extract test keys from the imported JSON (e.g., "1024x1024", "2048x512")
  const testKeys: string[] = Object.keys(tcpData);
  const [selectedTestKey, setSelectedTestKey] = useState<string>(testKeys[0]);

  // Get the time data for the selected test key
  const times: number[] = tcpData[selectedTestKey];

  // Convert each time measurement (in seconds) into throughput (MB/s)
  // Since each test sends 1 MB, throughput = 1 MB / time
  const throughputData: number[] = times.map((time) => 1 / time);

  // Prepare the chart data for the line chart
  const chartData = {
    labels: throughputData.map((_, index) => `Trial ${index + 1}`),
    datasets: [
      {
        label: `Throughput for ${selectedTestKey} (MB/s) ${file}`,
        data: throughputData,
        fill: false,
        borderColor: "rgba(255,99,132,1)",
        backgroundColor: "rgba(255,99,132,0.2)",
      },
    ],
  };

  // Chart options for responsiveness and styling
  const chartOptions = {
    responsive: true,
    plugins: {
      legend: { position: "top" as const },
      title: { display: true, text: `TCP Throughput Visualization : ${file}` },
    },
  };

  // Handler for changing the selected test key
  const handleTestKeyChange = (e: ChangeEvent<HTMLSelectElement>) => {
    setSelectedTestKey(e.target.value);
  };

  return (
    <div style={{ width: "80%", margin: "50px auto" }}>
      <h1>TCP Throughput Graph: {file}</h1>
      <div style={{ marginBottom: "20px" }}>
        <label>
          Select Test:&nbsp;
          <select value={selectedTestKey} onChange={handleTestKeyChange}>
            {testKeys.map((key) => (
              <option key={key} value={key}>
                {key}
              </option>
            ))}
          </select>
        </label>
      </div>
      <Line data={chartData} options={chartOptions} />
    </div>
  );
};

export default App;
