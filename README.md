# Sonrai

Sonrai provides tools to manage the process of importing and monitoring data. It provides a scheduling server, a handful of generally useful components
for scraping data, and a user interface for monitoring data imports. There are 3 primary concepts in Sonrai:

1. **Producers** are components that know how to connect to the original source of data and extract individual of pieces of data.
2. **Records** are created by producers and represent a single object containing the information the producers acquired. Examples include:
   end-of-day stock prices, measurements from a IoT sensor, or a single blog post. Tags are associated `record` with each record which are used
   to route data between different components in Sonrai.
3. **Transformers** operate on data records. They are used to enrich data records with additional information or to restructure data records.
4. **Syncs** perform actions based on data records--typically to save the record in a data lake.

## Getting Started

The sonrai server is built using mage. Build the server with the command `mage build` and run it with the command `sonrai serve`. This will start
the api service and provide access to the Web UI. Additional documentation is available in the `doc` directory.
