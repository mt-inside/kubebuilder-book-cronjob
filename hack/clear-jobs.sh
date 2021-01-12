for j in $(kubectl get jobs -o name)
do
    kubectl delete $j
done
