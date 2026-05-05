import { useEffect, useRef } from 'react';
import * as THREE from 'three';

type Props = { url: string };
export function PanoramaViewer({ url }: Props) {
  const mountRef = useRef<HTMLDivElement>(null);
  useEffect(() => {
    const mount = mountRef.current!;
    const scene = new THREE.Scene();
    const camera = new THREE.PerspectiveCamera(75, mount.clientWidth / mount.clientHeight, 0.1, 2000);
    const renderer = new THREE.WebGLRenderer({ antialias: true });
    renderer.setSize(mount.clientWidth, mount.clientHeight);
    mount.appendChild(renderer.domElement);
    const sphere = new THREE.Mesh(new THREE.SphereGeometry(500, 64, 64), new THREE.MeshBasicMaterial({ side: THREE.BackSide }));
    scene.add(sphere);
    const loader = new THREE.TextureLoader(); let texture: THREE.Texture | undefined;
    loader.load(url, (t) => { texture = t; (sphere.material as THREE.MeshBasicMaterial).map = t; (sphere.material as THREE.MeshBasicMaterial).needsUpdate = true; });
    let yaw=0,pitch=0,drag=false,lastX=0,lastY=0;
    const onDown=(e:MouseEvent)=>{drag=true;lastX=e.clientX;lastY=e.clientY};
    const onUp=()=>drag=false;
    const onMove=(e:MouseEvent)=>{ if(!drag)return; yaw += (e.clientX-lastX)*0.003; pitch += (e.clientY-lastY)*0.003; pitch=Math.max(-1.3,Math.min(1.3,pitch)); lastX=e.clientX; lastY=e.clientY; };
    const onWheel=(e:WheelEvent)=>{ camera.fov = Math.max(35, Math.min(100, camera.fov + e.deltaY * 0.03)); camera.updateProjectionMatrix(); };
    mount.addEventListener('mousedown',onDown); window.addEventListener('mouseup',onUp); window.addEventListener('mousemove',onMove); mount.addEventListener('wheel',onWheel);
    const anim=()=>{ camera.lookAt(new THREE.Vector3(Math.sin(yaw)*Math.cos(pitch), Math.sin(pitch), Math.cos(yaw)*Math.cos(pitch))); renderer.render(scene,camera); requestAnimationFrame(anim)}; anim();
    return ()=>{ texture?.dispose(); (sphere.material as THREE.Material).dispose(); sphere.geometry.dispose(); renderer.dispose(); mount.removeChild(renderer.domElement);} ;
  }, [url]);
  return <div className='h-full w-full' ref={mountRef} />;
}
