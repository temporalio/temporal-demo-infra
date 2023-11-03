# change order

1. PlaceOrder (execute workflow)
1. GetsOrderInformation

1. Requests authorization for prescription
1. Request fulfillment for prescription
1. While waiting both approval and fulfillment, change order comes in
1. This stops waiting, cancels in flight processes, and replays workflow with new prescription id
