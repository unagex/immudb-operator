# Immudb Configuration
After deploying the operator, you can create a `Immudb` resource to create a database ([example Immudb resource](../config/samples/v1_immudb.yaml)).

The following `spec values` can be updated:

| Name | Type | Default value | Description |
| --- | --- | --- | --- |
| image | string | "codenotary/immudb:latest" | The Immudb image. |
| imagePullPolicy | string | "IfNotPresent" | ImagePullPolicy of immudb image. |
| replicas | int | 1 | Number of replicas of immudb image. The value can only be 1 at the moment. The immudb team is working hard in adding replication in the future. |
| volume.storageClassName | string | Name of the default storageClass of your cluster.  | StorageClassName of the database. |
| volume.size | string | No default value, mandatory to set.  | Size of the database, e.g., 5Mi, 10Gi.  |