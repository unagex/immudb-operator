# Immudb Configuration
After deploying the operator, you can create a `Immudb` resource to create a database ([example Immudb resource](../config/samples/v1_immudb.yaml)).

The following `spec values` can be updated:

| name | type | default | meaning |
| --- | --- | --- | --- |
| image | string | "codenotary/immudb:latest" | Immudb image |
| imagePullPolicy | string | "IfNotPresent" | ImagePullPolicy of immudb image |
| replicas | int | nil (mandatory to set) | Number of replicas of immudb image. The value can only be 1 at the moment. The immudb team is working hard in adding replication in the future. |


 You can create many databases by creating multiple `Immudb`.


