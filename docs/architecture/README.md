# Architecture

Data is made available on ZGDA through the following flow:

1. The [Disperser](../spec/flows/disperer.md) encodes the data in accordance with the [storage module](../spec/flows/protocol-modules/storage/overview.md) requirements, constructs the appropriate header, and sends the chunks to the DA nodes.
2. Upon receiving signatures from the DA nodes, the disperser aggregates these signatures.
3. Next, the disperser sends the aggregated signatures and header to the `confirmBatch` method of the `ServiceManager`
4. Once retrievers see the confirmed the batch on chain, they can request to download the associated chunks from a set of ZeroG Storage Nodes, in accordance with the [retrieval module](../spec/flows/protocol-modules/retrieval/retrieval.md) of the protocol.

<figure><img src="../../.gitbook/assets/image (1).png" alt=""><figcaption><p>Figure 1. Architecture Overview </p></figcaption></figure>
