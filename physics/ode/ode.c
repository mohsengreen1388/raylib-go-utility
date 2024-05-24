#include <ode/ode.h>
#include <stdio.h>
#include "_cgo_export.h"
#include <raylib.h>
#include <raymath.h>

void callNearCallback(void *data, dGeomID obj1, dGeomID obj2)
{
	nearCallback(data, obj1, obj2);
}

void callMovedCallback(dBodyID body)
{
	movedCallback(body);
}

int callTriCallback(dGeomID mesh, dGeomID other, int index)
{
	return triCallback(mesh, other, index);
}

int callTriRayCallback(dGeomID mesh, dGeomID ray, int index, dReal u, dReal v)
{
	return triRayCallback(mesh, ray, index, u, v);
}
int x =0;
dTriMeshDataID TriMeshDataBuildSingle(Mesh plane, float scale)
{

	for (size_t i = 0; i < plane.vertexCount; i++)
	{
		Vector3 vec3 = Vector3Scale((Vector3){plane.vertices[i * 3 + 0], plane.vertices[i * 3 + 1], plane.vertices[i * 3 + 2]},scale);
		plane.vertices[i * 3 + 0] = vec3.x;
		plane.vertices[i * 3 + 1] = vec3.y;
		plane.vertices[i * 3 + 2] = vec3.z;
	}

	int nV = plane.vertexCount;
	int *groundInd = RL_MALLOC((nV) * sizeof(int));
	for (int i = 0; i < nV; i++)
	{
		groundInd[i] = i;
	}

	dTriMeshDataID triData = dGeomTriMeshDataCreate();

	dGeomTriMeshDataBuildSingle(triData, plane.vertices, 3 * sizeof(float), nV, groundInd, nV, 3 * sizeof(int));

	return triData;
}

dTriMeshDataID TriMeshDataBuildSingleAnm(Mesh plane)
{
	int nV = plane.vertexCount;

	int *groundInd = RL_MALLOC(nV * sizeof(int));
	for (int i = 0; i < nV; i++)
	{
		groundInd[i] = i;
	}

	dTriMeshDataID triData = dGeomTriMeshDataCreate();

	dGeomTriMeshDataBuildSingle1(triData,&plane.animVertices[0], 3 * sizeof(float), nV, groundInd, nV, 3 * sizeof(float),&plane.animNormals[0]);
	
	return triData;
}


