# wwapi

This is the warewulf API. Currently this is a work in progress. The basic plan is as follows:  

- The API server has nothing to do with how the warewulf CLI works externally.
- The API server should be a separate service and on a differernt port than warewulfd. warewulfd talks to the nodes.
- API server installation should be optional and not required.
- Existing code paths should call into the API. They need not call into the API server, but they should call API code to avoid divergence.
- For security, TLS makes sense.
- We need to change sync mutexes to flock. sync mutexes do not sync processes, just threads.
- We need to have a look at the os.Exit calls. Most of them should likely just return an error.


